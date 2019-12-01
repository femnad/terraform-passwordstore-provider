PLUGIN_DIR:=${HOME}/.terraform.d/plugins

all:
	mkdir -p ${PLUGIN_DIR}
	go build -o ${PLUGIN_DIR}
