# https://v2.vuejs.org/v2/cookbook/dockerize-vuejs-app.html?redirect=true#Real-World-Example
# build stage
FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
ENV NODE_ENV=production
ENV VUE_APP_API_URL=/api
# https://stackoverflow.com/questions/69394632/webpack-build-failing-with-err-ossl-evp-unsupported
ENV NODE_OPTIONS=--openssl-legacy-provider
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

# https://stackoverflow.com/questions/49955097/how-do-i-add-a-user-when-im-using-alpine-as-a-base-image
RUN apk update && apk upgrade && \
    apk add --no-cache curl && \
    addgroup -S -g 1001 www && \
    adduser -S -H -D -u 1001 -G www www && \
    chown -R www:www /var/cache/nginx && \
    chown -R www:www /var/log/nginx && \
    sudo -u www -g www touch /var/run/nginx.pid

EXPOSE 8080

CMD ["nginx", "-g", "daemon off;"]
