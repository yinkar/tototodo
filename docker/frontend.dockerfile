FROM node:16.13.0

WORKDIR /app

COPY frontend/package*.json ./

ADD frontend .

RUN npm install

CMD [ "npm", "run", "serve" ]