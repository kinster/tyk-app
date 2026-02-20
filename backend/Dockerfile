FROM node:20-alpine
WORKDIR /app
COPY server.js .
RUN npm init -y && npm install express
ENV PORT=3001
EXPOSE 3001
CMD ["node","server.js"]