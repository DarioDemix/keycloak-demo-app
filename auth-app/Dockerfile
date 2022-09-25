FROM node:current-alpine3.16 as BUILDER
WORKDIR /app

COPY package*.json .
RUN npm i

COPY index.html webpack* .
COPY src src
RUN npm run build

FROM nginx:stable
COPY --from=BUILDER /app/index.html /app/dist /usr/share/nginx/html