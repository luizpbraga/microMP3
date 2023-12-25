# Use the official MySQL image as the base image
FROM mysql:latest

# Environment variables to set MySQL configurations
ENV MYSQL_ROOT_PASSWORD=123
ENV MYSQL_DATABASE=auth
ENV MYSQL_USER=auth_user
ENV MYSQL_PASSWORD=auth
ENV MYSQL_PORT=3306
ENV MYSQL_HOST=localhost


# Create a directory to persist MySQL data
RUN mkdir -p /var/lib/mysql_data

# Set the working directory
WORKDIR /var/lib/mysql_data

# Define o volume para armazenar os dados do MySQL
VOLUME /var/lib/mysql_data

# Copy SQL script to initialize the database
COPY init.sql /docker-entrypoint-initdb.d/

# Grant permissions for the init script
RUN chmod -R 775 /docker-entrypoint-initdb.d/

# Expose the MySQL default port
EXPOSE 3306

# Start MySQL when the container launches
CMD ["mysqld"]
