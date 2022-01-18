FROM node:15.12.0

WORKDIR /app

COPY frontend/package*.json ./

ADD frontend .

RUN npm install

CMD [ "npm", "run", "dev" ]