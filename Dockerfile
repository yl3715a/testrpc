FROM alpine:3.22

# Define the project name | 定义项目名称
ARG PROJECT=restrpc
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=restrpc.yaml
# Define the author | 定义作者
ARG AUTHOR="example@example.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY ./rpc/${PROJECT}_rpc ./
COPY ./rpc/etc/${CONFIG_FILE} ./etc/

EXPOSE 8667

ENTRYPOINT ["./restrpc_rpc", "-f", "etc/restrpc.yaml"]