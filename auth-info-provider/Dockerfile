FROM node:current-alpine3.16

WORKDIR /app

COPY package* .
RUN npm i

COPY src .

CMD ["node", "index.js"]