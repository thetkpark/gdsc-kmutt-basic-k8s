FROM node:16-alpine as builder
WORKDIR /app
COPY ./package.json ./
COPY ./yarn.lock ./
RUN yarn
COPY . .
ENV REACT_APP_API_BASEURL=/api
RUN yarn build

FROM nginx
COPY --from=builder /app/build /usr/share/nginx/html