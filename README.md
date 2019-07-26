# Golang project structure

### The idea
It would be nice to have a structure that would fit all type of projects regardless 
how complicated the project is and regardless of what architectural style it uses 
(Hexagonal, n-tier, ...)

Is it possible or do we need to change the project structure always completely based on what we are 
building?

## Hypotheses 1
It's possible to have a basic general template and then leave out
out some directories or add nested directories based on the concrete project

## Hypotheses 2 - Counter Hypotheses for Hypotheses 1
There are too many different type of projects and it's important how big the project is to select
the appropriate directory structure.

# Approach
1. Create projects categorization
1. Collects examples directory structures
1. Try to combine them and create so called "master" directory structure that would combine all
1. ....
1. Profit :)


# Categorization of projects
## Type of projects
- Backend service with business logic with RPC or REST HTTP API
- API service with limited amount of business logic, more mapping and validation
- Library providing go methods as an API 
- Binary
- Web project serving HTML, JS, Images

## Size of the projects
- Small library providing a couple of methods
- A simple micro-service which provides key - value store via a REST API, has persistence
- Relatively small micro-service that has 1 Aggregate, Persistence, and communicates with an external service
- Middle size service that has 1 aggregate, persistence, REST API and RPC API and communicates with 
2 other services via RPC
- ... ?

## Architecture styles
- Hexagonal
- DDD - Domain Driven Design (might be close to hexagonal)
- 3-tier - Presentation, Business, Persistence
- Multilayer - Presentation, Application, Business, Persistence
- MVC - Model, View, Controller
- ... more ?

# Directories structures

## 1) Simple Lib
.... no main function

## 2) Hexagonal
    .
    ├── build                       # Compiled files (not commited, but appears in the project after first build)
    ├── contract                    # 
    │   ├── proto                   # 
    │   └── swagger                 # TODOdescribe
    ├── doc                         # Documentation (images, charts, more md files if needed)
    ├── gradle                      # Gradle wrapper
    ├── k8s                         # K8s deployment files
    ├── pkg                         # The actual source code of the service
    │   ├── config  
    │   ├── client  
    │   │   ├── externalservice.go 
    │   │   └-─ anothercompany.go 
    │   ├── handler    
    │   ├── domain
    │   │   ├── user.go 
    │   │   ├── userid.go 
    │   │   └── address.go 
    │   └── repository
    │       ├── usermemory.go 
    │       └── userpostgres.go
    ├── .gitignore
    ├── docker-compose.yml
    ├── Jenkinsfile
    ├── main.go                       # Main file of the service
    ├── Makefile
    ├── LICENSE
    └── README.md

## 3) Hexagonal for 3 Aggregates (DDD terminology)
The question is if such a service would make sense ... mixing users, invoices and products. 

For the beginning probably yes. When we don't know to much about the domain and we have only 1 team of 
people the development will be probably faster with 1 service.
And if in the future when the service grows and the teams split into multiple teams, 
we can split also the service into more.

    .
    ├── build                 # Compiled files (not commited, but appears in the project after first build)
    ├── contract              # 
    │   ├── proto             # 
    │   └── swagger           # 
    ├── doc                   # Documentation (images, charts, more md files if needed)
    ├── gradle                # Gradle wrapper
    ├── k8s                   # K8s deployment files
    ├── pkg                   # The actual source code of the service
    │   ├── config   
    │   ├── app               # Application services (responsibility: auth?, transactions, delegation of logic to domain, delegation to persistence)
    │   ├── handler           # If we don't have too many files we might not needs subdirs here
    │   │   ├── user.go
    │   │   ├── invoice.go
    │   │   ├── product.go
    │   ├── domain            # I guess here we should have more sub-directories here not to get all things mixed?
    │   │   ├── user          # I'm not sure about this...
    │   │   ├── invoice 
    │   │   └── product 
    │   └── repository # How many things will be here? do we need separation of repos for different aggregates?
    ├── .gitignore
    ├── docker-compose.yml
    ├── Jenkinsfile
    ├── main.go               # Main file of the service
    ├── Makefile
    ├── LICENSE
    └── README.md
    
## 4) N-Tier
    .
    ├── build                 # Compiled files (not commited, but appears in the project after first build)
    ├── contract              # 
    │   ├── proto             # 
    │   └── swagger           # 
    ├── doc                   # Documentation (images, charts, more md files if needed)
    ├── gradle                # Gradle wrapper
    ├── k8s                   # K8s deployment files
    ├── pkg                   # The actual source code of the service
    │   ├── application   
    │   ├── config   
    │   ├── controller    
    │   ├── entities   
    │   └── repository
    ├── .gitignore
    ├── docker-compose.yml
    ├── Jenkinsfile
    ├── main.go               # Main file of the service
    ├── Makefile
    ├── LICENSE
    └── README.md


    
    
## 4) More binaries
.... with the `cmd` folder


# Results (in progress)

## Hypothesis 2 confirmation
- This repository https://github.com/golang-standards/project-layout has quite some many (8000 atm) which 
suggests that many people like it.


# Resources
- https://github.com/golang-standards/project-layout
- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- https://rakyll.org/style-packages/
- https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2
- https://www.ardanlabs.com/blog/2017/02/design-philosophy-on-packaging.html
- https://peter.bourgon.org/go-best-practices-2016/#repository-structure
- https://vsupalov.com/go-folder-structure/
- https://en.wikipedia.org/wiki/List_of_software_architecture_styles_and_patterns