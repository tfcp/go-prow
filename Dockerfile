FROM registry-vpc.cn-hangzhou.aliyuncs.com/medlinker/alpine:3.8

LABEL maintainer="zhaosuji@medlinker.com"

#######################################################
# 同时适用于单仓库应用和大仓库应用
#######################################################

# 设置环境变量
ENV APP_NAME project

ENV APP_ROOT /var/www
ENV APP_PATH $APP_ROOT/$APP_NAME
ENV CONFIG_PATH $APP_PATH/config
ENV LOG_ROOT /var/log/project
ENV LOG_PATH /var/log/project/$APP_NAME
ENV PATH $APP_PATH/scripts:$PATH
ENV CONFIG_CENTER_BASE_URL http://consul.infra.svc.cluster.local:8500

# 创建配置目录
RUN mkdir -p $CONFIG_PATH
# 创建日志目录
RUN mkdir -p $LOG_PATH && chmod 777 -R $LOG_PATH

# 执行入口文件添加
ADD ./main $APP_PATH/
ADD ./scripts/*.sh $APP_PATH/scripts/
RUN chmod +x $APP_PATH/scripts/*.sh

# 启动
CMD ["start.sh"]
