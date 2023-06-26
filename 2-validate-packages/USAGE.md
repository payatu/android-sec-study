# Build play-head

$ go build

# Build FireProx docker image (Docker must be installed)

$ git clone https://github.com/ustayready/fireprox 
$ cd fireprox
$ docker build -t fireprox .

# Generate AWS API Gateway

$ docker run --rm -it fireprox --command create --access_key $AWS_ACCESS_KEY --secret_access_key $AWS_SECRET_KEY --region us-east-2 --url "https://play.google.com/store/apps/details"

# Populate the ./endpoints.txt file with generated AWS API Gateway hostnames

$ echo "*.execute-api.us-east-2.amazonaws.com" > endpoints.txt

# Run play-head

$ ./play-head list-of-package-names.txt
