# Greenlight API

## Development

For local development, we have to start the postgres server and export the login information.
After we have done this, we can run make run to build and run the aplication.

1. `export GREENLIGHT_DB_DSN=postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]`
1. `export GREENLIGHT_SMTP_USERNAME=<username>`
1. `export GREENLIGHT_SMTP_PASSWORD=<password>`
1. `sudo -u postgres systemctl start postgresql.service`
1. `make run`


