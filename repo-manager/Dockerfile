FROM alpine:3.1

# Update
RUN apk add --update python py-pip

# Create app directory
WORKDIR /usr/src/app

# Copying the requirements.txt file
COPY requirements.txt requirements.txt

# Install app dependencies
RUN pip install -r requirements.txt

# Bundle app source
COPY simpleapp.py simpleapp.py

EXPOSE  8000
CMD ["python", "simpleapp.py", "-p 8000"]
