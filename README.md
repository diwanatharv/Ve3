This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `npm run build`

Builds the app for production to the `build` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br />


**Server: Golang
Database: Local MongoDB,Redis**
DB connection string, name and collection name moved to `.env` file as environment variable.
golang https://golang.org/dl/
2. echo package for router `go get github.com/labstack/echo`
3. mongo-driver package to connect with mongoDB `go get go.mongodb.org/mongo-driver`
## :computer: Start the application

1. Make sure your mongoDB is started
2. use `go mod tidy` to download all the dependencies
3.  From go-server directory, open a terminal and run
    `go run main.go`
4. From client directory,  
   a. install all the dependencies using `npm install`  
   b. start client `npm start`