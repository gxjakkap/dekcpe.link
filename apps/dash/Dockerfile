# syntax=docker.io/docker/dockerfile:1

FROM oven/bun:alpine AS base
RUN apk update && apk add --no-cache curl

FROM base AS deps
RUN apk add --no-cache gcompat
WORKDIR /app

COPY package.json bun.lock* ./
COPY apps/dash/package.json ./apps/dash/
COPY packages/db/package.json ./packages/db/
COPY packages/const/package.json ./packages/const/
COPY apps/api/package.json ./apps/api/

COPY packages/db ./packages/db
COPY packages/const ./packages/const

RUN bun install --frozen-lockfile

FROM base AS builder
WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Build Args
ARG AUTH_SECRET
ARG AUTH_URL
ARG NEXT_PUBLIC_BASE_URL
ARG NEXT_PUBLIC_SHORTLINK_BASE_URL

ENV NEXT_TELEMETRY_DISABLED=1
ENV NODE_ENV=production
ENV AUTH_SECRET=$AUTH_SECRET
ENV AUTH_URL=$AUTH_URL
ENV NEXT_PUBLIC_BASE_URL=$NEXT_PUBLIC_BASE_URL
ENV NEXT_PUBLIC_SHORTLINK_BASE_URL=$NEXT_PUBLIC_SHORTLINK_BASE_URL

ENV PATH="/app/node_modules/.bin:$PATH"

WORKDIR /app/apps/dash
RUN bun run build

FROM base AS runner
WORKDIR /app

ENV NODE_ENV=production
ENV PORT=4000
ENV HOSTNAME="0.0.0.0"

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs


COPY --from=builder /app/apps/dash/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/apps/dash/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/apps/dash/.next/static ./apps/dash/.next/static

USER nextjs
EXPOSE 4000

CMD ["bun", "run", "apps/dash/server.js"]
