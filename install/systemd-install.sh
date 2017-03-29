#!/usr/bin/env bash

# TODO SEEDMS change seedms the microservice's name (prefer it similar to name constant in main.go)
NAME="seedms"
BUILD_NAME="app"
CONF_DIR="/etc/${NAME}"
CONF_FILE="${CONF_DIR}/${NAME}.conf.yaml"
INSTALL_DIR="/usr/local/bin"
UNIT_FILE="/etc/systemd/system/${NAME}.service"
EXIT_CODE_FAIL=1

./systemd-uninstall.sh || exit ${EXIT_CODE_FAIL}
echo "Begin install"
mkdir -p "${CONF_DIR}" || exit ${EXIT_CODE_FAIL}
if [ ! -f "${CONF_FILE}" ]; then
    cp "conf.yaml" "${CONF_FILE}" || exit ${EXIT_CODE_FAIL}
fi
mkdir -p "${INSTALL_DIR}" || exit ${EXIT_CODE_FAIL}
cp -f "${BUILD_NAME}" "${INSTALL_DIR}/${NAME}" || exit ${EXIT_CODE_FAIL}
cp -f "unit.service" "${UNIT_FILE}" || exit ${EXIT_CODE_FAIL}
systemctl enable "${NAME}.service"
echo "Config file is at '${CONF_FILE}'"
echo "Install complete"
