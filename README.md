# go-auth
Belajar Authentication & Authorization with Go MariaDB

# Framework Libraries
    echo


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

# CI continuous integration
## Docker File
## GitHub Actions
## token github
    git push https://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@github.com/absormu/go-auth.git