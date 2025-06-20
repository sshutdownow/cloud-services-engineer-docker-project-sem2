# https://v2.vuejs.org/v2/cookbook/dockerize-vuejs-app.html?redirect=true#Real-World-Example
# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
ENV NODE_ENV=production
ENV VUE_APP_API_URL=/api
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

RUN apk update && apk upgrade && \
    apk add --no-cache curl && \
    addgroup -g 1001 -S www && \
    adduser -S -D -u 1001 -G www www && \
    chown -R www:www /var/cache/nginx /var/log/nginx \
    touch /var/run/nginx.pid && \
    chown -R webuser:webgroup /var/run/nginx.pid

EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
