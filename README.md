FeatureFlagManagerService is a manager of feature-flags. This project was the subject of my intership in the compagny Citron!
I thanl them for all that I have learned.

TO CREATE AND INIT A DATA BASE FOR FEATURE FLAGS (for testing API)
To create a database:

docker run --name container_flags -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=docker -e POSTGRES_DB=test_feature_flag -p 5432:5432 -d postgres

YOU HAVE TO CONNECT TO YOUR DOCKER CONTAINER, THEN: psql -h localhost -U postgres

AND TO CONNECT TO THE DATABASE: \c feature_flag

THEN copy the following lignes in your data base to init it:

CREATE TABLE feature_flags ( Id serial, slug VARCHAR(50), Label VARCHAR(50), isEnabled BOOL, Application VARCHAR(50), Projects VARCHAR(50), Owners VARCHAR(50), Description VARCHAR(50), CreatedAt TIMESTAMP WITH TIME ZONE, UpdatedAt TIMESTAMP WITH TIME ZONE, PRIMARY KEY (slug, application) );

IF YOU WANT TO CREATE A NEW TABLE FOR THE APPLICATIONS:

CREATE TABLE applications ( Id serial, Label VARCHAR(50), Description VARCHAR(50), CreatedAt TIMESTAMP WITH TIME ZONE, UpdatedAt TIMESTAMP WITH TIME ZONE, PRIMARY KEY (label) );

TO EXECUTE TEST ON DB
1) Create your db docker run --name container_flags -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=docker -e POSTGRES_DB=postgres -p 5432:5432 -d postgres

2) In your terminal: go test -v ./...

Config and environment variables to run programm locally
1)Create on docker your db, so copy this lines in your: docker run -d \ --name my_postgres \ -p 5432:5432 \ -v $(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql \ -e POSTGRES_DB=feature_flag \ -e POSTGRES_USER=postgres \ -e POSTGRES_PASSWORD=docker \ postgres:latest

3)In your terminal, copy the following lines: export DATABASE_HOST=localhost && export DATABASE_USER=postgres && export DATABASE_PASS=docker && export DATABASE_NAME=feature_flag && export CONFIG_JWT_PRIVATE_KEY=citron_c_est_super_! && go run cmd/app/main.go

Get your token
To have access to featureflago, you must have authorization. How? Thanks to your token in the headers of your request, with the name "autorization". We use JWT Token : your token is divided in 3 parts: one for the algorithm which crypt your data, one for you data (email, name, firstname... ) and the last one for the security key. To create one with your own information: 1/ go on https://jwt.io/,



Starting
Then, you're set!