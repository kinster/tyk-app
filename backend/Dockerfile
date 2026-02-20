FROM node:20-alpine

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm install

COPY server.js .

ENV PORT=3001
EXPOSE 3001

CMD ["node", "server.js"]