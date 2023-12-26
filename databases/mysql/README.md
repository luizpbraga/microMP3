The auth mysql database stores data relate to the authentication system.

### usar o root

```
docker exec -it mysql-container mysql -u root -p
```

### builda

```
docker build -t mysql-auth .
```

### roda

```
docker run -d --name mysql-auth-container -v ${CWD}/data:/var/lib/mysql_data -p 3308:3360 mysql-auth:latest
```

### para

```
docker stop mysql-auth-container
```

### re(iniciat)

```
    docker start mysql-auth-container
```
