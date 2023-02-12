# Build Stage
FROM node:19.6.0-alpine3.17 AS builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN npm install && npm run build

# Deploy Stage
FROM nginx:mainline-alpine-slim
COPY ./docker/nginx.conf /etc/nginx/nginx.conf
# Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 3000 80
ENTRYPOINT ["nginx", "-g", "daemon off;"]
# EOF