CREATE TABLE palabras
(
    palabra varchar(255) COLLATE utf8mb4_spanish_ci,
    tipo varchar(255),
    genero varchar(255),
    numero varchar(255),
    raiz varchar(255),
    afijo varchar(255),
    tonica integer,
    silabas varchar(255),
    locale varchar(255),
    origen varchar(255),
    sinonimos MEDIUMTEXT
);

LOAD DATA INFILE '/opt/palabras.csv' 
INTO TABLE palabras
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;
