FROM node:16-alpine as builder

WORKDIR /app

RUN corepack enable

COPY package.json yarn.lock .yarnrc.yml .pnp.cjs .pnp.loader.mjs  ./
COPY .yarn/cache/ ./.yarn/cache/
COPY .yarn/releases/ ./.yarn/releases/

RUN yarn install --immutable --immutable-cache

COPY . ./

RUN yarn build


FROM nginx:stable-alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /app/build/ ./
