# userrestapigo
userrestapigo golang using rest api, mysql 

# used packages
go get github.com/go-sql-driver/mysql
# env loader
go get github.com/joho/godotenv
# HTTP request router and URL matcher for Go
go get github.com/gorilla/mux
# validations 
go get github.com/go-playground/validator/v10

# run database
mysql -u root -p < database/schema.sql