# ─── Build Stage ─────────────────────────────────────────────────────────────
FROM node:24-slim AS builder

ARG VITE_API_URL=http://localhost:8080
ARG VITE_WS_URL=ws://localhost:8080

ENV VITE_API_URL=$VITE_API_URL
ENV VITE_WS_URL=$VITE_WS_URL

WORKDIR /app

COPY package*.json ./

RUN env

RUN npm install

COPY . .

RUN npm run build

# ─── Runtime Stage ───────────────────────────────────────────────────────────
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080

HEALTHCHECK --interval=10s --timeout=3s --retries=3 \
    CMD wget --quiet --tries=1 --spider http://localhost/ || exit 1

CMD ["nginx", "-g", "daemon off;"]
