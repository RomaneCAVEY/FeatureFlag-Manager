
create table if not exists feature_flags (
Id  serial,
slug VARCHAR(200),
Label VARCHAR(200),
isEnabled BOOL,
Application VARCHAR(200),
Projects VARCHAR(200),
Owners VARCHAR(200),
description VARCHAR(200),
CreatedAt TIMESTAMP WITH TIME ZONE,
UpdatedAt TIMESTAMP WITH TIME ZONE,
CreatedBy VARCHAR(200),
UpdatedBy VARCHAR(200),
PRIMARY KEY (slug, application)
);


create table if not exists applications (
Id  serial,
Label VARCHAR(200),
Description VARCHAR(200),
CreatedAt TIMESTAMP WITH TIME ZONE,
UpdatedAt TIMESTAMP WITH TIME ZONE,
PRIMARY KEY (label)
);
