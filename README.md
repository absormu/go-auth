# go-auth
Belajar Authentication & Authorization with Go MariaDB

# Framework Libraries
    echo
    go get -u github.com/go-sql-driver/mysql
    go get golang.org/x/crypto/bcrypt


# Docker 
### Setup Database
    nama database : dbauth
    username : goauth
    password : THTqAOELuFckJZZaBP7Z
    jenis database : MariaDB

### Db Container
    mkdir dbdata
    docker run -itd --name=dbauth -p 3307:3306 -v $(pwd)/dbdata:/var/lib/mysql --env="MYSQL_ROOT_PASSWORD=root212" mariadb:10.5
    docker exec -it dbauth bash
    mysql -u root -proot212

    create database dbauth;
    use dbauth;
    CREATE USER 'goauth'@'%' IDENTIFIED BY 'THTqAOELuFckJZZaBP7Z';
    GRANT ALL PRIVILEGES ON *.* TO 'goauth'@'%';  
    FLUSH privileges;

    name container : dbauth
    docker container start dbauth
    docker inspect dbauth
    docker container stop dbauth
    docker rm dbauth

### Table user
    create table user(
        id BIGINT(20) NOT NULL AUTO_INCREMENT,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(150) NOT NULL,
        password VARCHAR(150) NOT NULL,
        user_contact_id BIGINT(20) NOT NULL,
        role_id INT(11) NOT NULL,
        active tinyint(1) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(100) NOT NULL,
        modified_at TIMESTAMP NULL, 
        modified_by VARCHAR(100),
        is_deleted tinyint(1) DEFAULT 0 NOT NULL,
        PRIMARY KEY ( id )
    );

    INSERT INTO `user` (name,email,password,user_contact_id,role_id,active,created_at,created_by,modified_at,modified_by,is_deleted) VALUES
	 ('Absor MU','admin@absormu.id','U21AcnREZWwhdmVyeQ==',1,1,1,'2022-02-21 15:37:52.0','SYSTEM',NULL,NULL,0);
### Table user_contact
    create table user_contact(
        id BIGINT(20) NOT NULL AUTO_INCREMENT,
        name VARCHAR(100) NOT NULL,
        city_id INT(11) NOT NULL,
        telephone VARCHAR(30) NULL,
        address TEXT NULL,
        type tinyint(1) NOT NULL, 
        active tinyint(1) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        created_by VARCHAR(100) NOT NULL,
        modified_at TIMESTAMP NULL, 
        modified_by VARCHAR(100),
        is_deleted tinyint(1) DEFAULT 0 NOT NULL,
        PRIMARY KEY ( id )
    );


# CI continuous integration
## Docker File
## GitHub Actions
## token github
    git push https://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@github.com/absormu/go-auth.git
