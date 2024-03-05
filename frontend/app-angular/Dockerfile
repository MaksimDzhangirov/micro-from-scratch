# этап сборки (build stage)
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install -g @angular/cli
RUN npm install
COPY . .
#RUN ng build

# этап production (production-stage)
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
