#!/bin/bash

PROJECT_NAME="bdntoy"
VERSION=v0.0.2
RELEASE_NAME="${PROJECT_NAME}"-"${VERSION}"
RELEASE_BODY="Release version ${VERSION}"
OUTPUT_DIR="bin" # 打包目录

# 平台和架构数组
PLATFORMS=("linux" "windows" "darwin")
ARCHITECTURES=("amd64" "arm64")

# 编译不同平台的应用
for PLATFORM in "${PLATFORMS[@]}"
do
    for ARCH in "${ARCHITECTURES[@]}"
    do
        echo "Building for ${PLATFORM} ${ARCH}..."
        GOOS=${PLATFORM} GOARCH=${ARCH} go build -o ${OUTPUT_DIR}/${PROJECT_NAME}_${PLATFORM}_${ARCH}
    done
done

# 创建一个压缩文件
echo "Creating release package..."
RELEASE_TAR="${RELEASE_NAME}.tar.gz"
tar -czvf "${RELEASE_TAR}" -C ${OUTPUT_DIR} .

RELEASE_ZIP="${RELEASE_NAME}.zip"
zip -r "${RELEASE_ZIP}" ${OUTPUT_DIR}/*

# 清理临时文件
echo "Cleaning up..."
rm -rf ${OUTPUT_DIR}
#rm "${RELEASE_TAR}"
#rm "${RELEASE_ZIP}"

echo "Release process completed!"