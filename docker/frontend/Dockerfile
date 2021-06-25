FROM node:16-alpine3.11
RUN mkdir /myapp
WORKDIR /myapp
RUN yarn add nuxt
COPY . /myapp

# Start the main process.
CMD ["yarn", "dev"]
