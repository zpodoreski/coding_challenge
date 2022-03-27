# coding_challenge

Teltech Backend Coding Challenge #1

Specification
- For this project, you will develop a simple web application in Go which accepts math problems via the URL and returns the response in JSON
- The application should be simple to test using curl or wget
- Example:
  http://localhost/add?x=2&y=5
  Output:
  {“action”: “add”, “x”: 2, “y”: 5, “answer”, 7, “cached”: false}
- Implement add, subtract, multiply, and divide operations
- Only two arguments will ever be passed: x and y (no need to handle a variable number of arguments)
- Cache results so that repeated calls with the same problem will return the answer from the cache
 - Show in the output JSON whether the cache was used
 - Expire anything in the cache that has not been hit for one minute

Additional notes and requirements
- The code should be delivered via Github (if you choose to create a private repo, add proper - - - access for the reviewer)
- Use your judgment in handling anything not explicitly defined in the requirements
- The result should be simple but robust
 - We will be looking for some things which are not specifically mentioned in order to gauge your - experience and knowledge of best practices
- Mistakes due to an obvious lack of familiarity with Go will not be judged harshly
- Describe how would you deploy your application (in the cloud) based on your experience
- In this and all other coding projects, readability is key
 - When in doubt, choose clarity over cleverness
​- Assume that we will try to break your app

Final notes
- Projects without a README.md file will be ignored
- After writing and testing your code, take a break and code-review it hours later, or the next day. Better to proofread before the editor gets his hands on it.
- Clarification questions are welcome and encouraged


README
 - This web application is a representation of a simple calculator that listen on specific path for    specific operation on two parameters (x and y)

 - Prior starting the program, Redis database is needed for program to work
 - To start Redis database there are two ways:
   - on Mac OSX install Docker, open it via launchpad then run script initRedisOSX.sh by executing command ./initRedisOSX.sh in root directory of project
   - that script is starting Docker container with Redis running on designated port

   - on Linux launch script initRedisUnix.sh by executing command ./initRedisUnix.sh in root directory of project
   - that script installs Docker and pull Redis image and start Docker container with Redis running on designated port

 - Web application can be started locally in two ways:
  - by running command: make build && cd bin && ./api
    - this command builds the program, enters directory where executable is and start it

  - or by: cd cmd && go run main.go
    - this command moves itself to cmd directory and starts locally go program

 - Web application listen on:
 http://localhost/{name of operation}?x={value of param x as number}&y={value of param y as number}

few examples:
  http://localhost/add?x=2&y=3

  http://localhost/divide?x=6&y=2

  http://localhost/multiply?x=3&y=2

  http://localhost/subtract?x=6&y=2


 - Response is always returned formatted in JSON with few params:
  Example:
  {"action":"add","answer":3,"cached":true,"x":1,"y":2}

  {"action":"divide","answer":3,"cached":false,"x":6,"y":2}

  {"action":"multiply","answer":2,"cached":false,"x":1,"y":2}

  {"action":"subtract","answer":2,"cached":true,"x":3,"y":1}

 - Cached value means if result was previously cached and that result was returned from cache
 - Result is cached for one minute, if not hit in that time, result is invalidated
 - When cache is hitted result is returned from cache and sets new cache with 1 minute ttl

 Cloud deployment
 - Operations are divided in functions because each operation represents one function
 - Every function uses shared code in pkg folder
 - Deployment is done with deploy.sh script
 - Deploy.sh script is copying pkg directory into the function root directory that is being deployed, after that is renaming imports to look inside local pkg folder instead of shared one, and initializing go.mod and go.sum to handle import
  - after deployment is done, pkg folder, go.mod, go.sum are being deleted from function and imports are      renamed back to look in root pkg instead of local

  - This way code can be tested and runned locally with one pkg folder to change, while deployment for google cloud works the way is should and the application runs on GCP





