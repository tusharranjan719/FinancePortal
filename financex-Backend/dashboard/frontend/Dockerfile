FROM node:12.16.3-alpine3.11

ARG VUE_CLI_MODE=production

WORKDIR /frontend

COPY package*.json ./

RUN npm install

COPY . .

RUN ./node_modules/.bin/vue-cli-service build --mode ${VUE_CLI_MODE}

EXPOSE 8080

CMD [ "npx", "serve", "-l", "tcp://0.0.0.0:8080", "dist" ]
