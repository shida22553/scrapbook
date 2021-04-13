FROM node:14.4.0
RUN mkdir /myapp
WORKDIR /myapp
RUN yarn add nuxt
COPY . /myapp

# Start the main process.
CMD ["yarn", "dev"]
