# Build fetch-meta

$ go get fetch-meta
$ go build

# Build FireProx docker image (Docker must be installed)

$ git clone https://github.com/ustayready/fireprox 
$ cd fireprox
$ docker build -t fireprox .

# Generate AWS API Gateway

$ docker run --rm -it fireprox --command create --access_key $AWS_ACCESS_KEY --secret_access_key $AWS_SECRET_KEY --region us-east-2 --url "https://play.google.com/store/apps/details"

# Populate the /tmp/endpoints.txt file with generated AWS API Gateway hostnames

$ echo "*.execute-api.us-east-2.amazonaws.com" > /tmp/endpoints.txt

# Run play-head

$ ./fetch-meta list-of-validated-package-names.txt
