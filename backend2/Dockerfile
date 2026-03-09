FROM node:20-alpine

WORKDIR /app

COPY package.json package-lock.json* ./
RUN npm install

COPY server.js .

ENV PORT=3002
EXPOSE 3002

CMD ["node", "server.js"]
