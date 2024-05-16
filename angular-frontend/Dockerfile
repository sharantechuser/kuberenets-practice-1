# FROM node:latest
# WORKDIR /app
# RUN npm install -g @angular/cli@13
# COPY package.json package-lock.json ./
# RUN npm ci
# COPY . .
# EXPOSE 4200
# CMD ["ng", "serve", "--host", "0.0.0.0"]
# Stage 1: Build the Angular application
FROM node:18.20.2-alpine3.19 AS build

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build

# Stage 2: Serve the Angular application using nginx
FROM nginx:alpine

COPY --from=build /app/dist/angular-frontend/browser /usr/share/nginx/html

EXPOSE 80

# CMD ["nginx", "-g", "daemon off;"]