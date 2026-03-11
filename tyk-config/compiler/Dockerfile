FROM tykio/tyk-plugin-compiler:v5.11.0

RUN go install github.com/TykTechnologies/tyk-sync@latest
ENV PATH=/root/go/bin:$PATH

COPY middleware/ /plugins/
COPY compiler/compile-all.sh /compile-all.sh
RUN chmod +x /compile-all.sh

ENTRYPOINT ["/compile-all.sh"]
