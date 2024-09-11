package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/dto"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/service"
	"github.com/RomaneCAVEY/FeatureFlag-Manager/infrastructure"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var errorFormatRequestFeatureFlag = errors.New("Wrong Format Request. The right format for request is: { label: string,\n application: string, \n isEnabled: bool,\n owners: []string, \n description:string (not necessary )\n ,projects: []string (not necessary)")
var errorFormatRequestApplication = errors.New("Wrong Format Request. The right format for request is: { label: string,\n description:string}")
var errorNoUser = errors.New("No User defined")

func main() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"X-Total-Count"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	infrastructure.ParamConnection()
	privateKey := infrastructure.GetHMACSecret()

	featureFlagRepository := infrastructure.FeatureFlagRepository{Collection: infrastructure.Connect()}
	ApplicationRepository := infrastructure.ApplicationRepository{Collection: infrastructure.Connect()}
	managerServiceApplication := service.MakeApplicationManagerService(&ApplicationRepository)
	managerServiceFeatureFlag := service.MakeFeatureFlagManagerService(&featureFlagRepository, &ApplicationRepository)

	r := gin.Default()
	r.Use(cors.New(config))
	r.GET("/flags/:id", func(c *gin.Context) {

		requiredId, _ := strconv.Atoi(c.Param("id"))
		GetFeatureFlagById, err := managerServiceFeatureFlag.GetFeatureFlagsById(uint32(requiredId))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Writer.Header().Set("X-Total-Count", "10")
		c.JSON(http.StatusOK, GetFeatureFlagById)

	})
	r.GET("/flags", func(c *gin.Context) {
		start := c.DefaultQuery("_start", "0")
		end := c.DefaultQuery("_end", "100")
		start_int, _ := strconv.Atoi(start)
		end_int, _ := strconv.Atoi(end)
		application := c.DefaultQuery("application", "all")

		if application == "all" {
			listAllFeatureFlags, count, err := managerServiceFeatureFlag.GetAllFeatureFlags(start_int, end_int)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Writer.Header().Set("X-Total-Count", strconv.Itoa(count))
			c.JSON(http.StatusOK, listAllFeatureFlags)

		} else {
			listAllFeatureFlagsByApplication, count, err := managerServiceFeatureFlag.GetFeatureFlagsByApplication(application, start_int, end_int)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.Writer.Header().Set("X-Total-Count", strconv.Itoa(count))
			c.JSON(http.StatusOK, listAllFeatureFlagsByApplication)
		}
	})

	r.Use(infrastructure.Authenticate(privateKey))


	r.PUT("/flags/:id", func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": error.Error(errorNoUser)})
			return
		}
		User := user.(entities.User)
		err := infrastructure.ValidateRequestFromCompagnyUser(User)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var json dto.ModifyFeatureFlagDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			fmt.Print(json)
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requiredId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ModifiedFeatureFlag, err := managerServiceFeatureFlag.ModifyFeatureFlag(uint32(requiredId), json, User)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Writer.Header().Set("X-Total-Count", "10")
		c.JSON(http.StatusOK, ModifiedFeatureFlag)

	})

	r.POST("/flags", func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": error.Error(errorNoUser)})
			return
		}
		User := user.(entities.User)

		err := infrastructure.ValidateRequestFromCompagnyUser(User)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		var json dto.CreateAFeatureFlagDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		featureFlagCreated, error := managerServiceFeatureFlag.CreateAFeatureFlag(json, User)

		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
		c.Writer.Header().Set("X-Total-Count", "10")
		c.JSON(http.StatusCreated, featureFlagCreated)
	})

	r.DELETE("/flags/:id", func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": error.Error(errorNoUser)})
			return
		}
		User := user.(entities.User)
		err := infrastructure.ValidateRequestFromCompagnyUser(User)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		requiredId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		error := managerServiceFeatureFlag.DeleteFeatureFlag(uint32(requiredId))
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
	})

	r.POST("/applications", func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": error.Error(errorNoUser)})
			return
		}
		User := user.(entities.User)
		err := infrastructure.ValidateRequestFromCompagnyUser(User)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		var json dto.CreateAnApplicationDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		applicationCreated, error := managerServiceApplication.CreateAnApplication(json)

		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
		c.Writer.Header().Set("X-Total-Count", "10")
		c.JSON(http.StatusCreated, applicationCreated)
	})

	r.GET("/applications", func(c *gin.Context) {
		start := c.DefaultQuery("_start", "0")
		end := c.DefaultQuery("_end", "100")
		start_int, _ := strconv.Atoi(start)
		end_int, _ := strconv.Atoi(end)
		listAllApplications, count, err := managerServiceApplication.GetAllApplications(start_int, end_int)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Writer.Header().Set("X-Total-Count", strconv.Itoa(count))
		c.JSON(http.StatusOK, listAllApplications)

	})
	r.GET("/applications/:id", func(c *gin.Context) {

		requiredId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		applicationById, err := managerServiceApplication.GetApplicationById(uint32(requiredId))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Writer.Header().Set("X-Total-Count", "10")
		c.JSON(http.StatusOK, applicationById)

	})

	r.PUT("/applications/:id", func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": error.Error(errorNoUser)})
			return
		}
		User := user.(entities.User)
		err := infrastructure.ValidateRequestFromCompagnyUser(User)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		var json dto.ModifyApplicationDTO
		if err := c.ShouldBindJSON(&json); err != nil {
			fmt.Print(json)
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requiredId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ModifiedApplication, err := managerServiceApplication.ModifyApplication(uint32(requiredId), json)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Writer.Header().Set("X-Total-Count", "10")
		c.JSON(http.StatusOK, ModifiedApplication)

	})

	r.DELETE("/applications/:id", func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": error.Error(errorNoUser)})
			return
		}
		User := user.(entities.User)
		err := infrastructure.ValidateRequestFromCompagnyUser(User)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		requiredId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		error := managerServiceApplication.DeleteApplication(uint32(requiredId))
		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
			return
		}
	})



	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8099"
	}
	r.Run(":" + httpPort)
}
