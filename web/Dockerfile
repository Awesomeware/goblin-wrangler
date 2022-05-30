FROM node:lts-alpine

ARG VITE_API_BASE_URL
ENV VITE_API_BASE_URL ${VITE_API_BASE_URL}

# install simple http server for serving static content
RUN npm install -g http-server

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./

# install project dependencies
RUN npm install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# build
RUN npm run build:dev

CMD [ "http-server", "dist" ]