FROM buoyantio/emojivoto-svc-base:graphql

ARG svc_name

COPY $svc_name/target/ /usr/local/bin/
RUN /usr/bin/sqlite3 /usr/local/bin/temp.db

# ARG variables arent available for ENTRYPOINT
ENV SVC_NAME $svc_name
ENTRYPOINT cd /usr/local/bin && $SVC_NAME
