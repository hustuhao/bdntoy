#!/bin/bash

PROJECT_NAME="bdntoy"
VERSION=v0.0.3
RELEASE_NAME="${PROJECT_NAME}-${VERSION}"
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
        # 创建输出目录
        PLATFORM_OUTPUT_DIR="${OUTPUT_DIR}/${PLATFORM}_${ARCH}"
        mkdir -p ${PLATFORM_OUTPUT_DIR}

        # windows 平台需要加后缀.exe
        if [ ${PLATFORM} = "windows" ]; then
            OUTPUT_NAME="${PROJECT_NAME}.exe"
        else
            OUTPUT_NAME="${PROJECT_NAME}"
        fi

        # 设置环境变量并构建
        GOOS=${PLATFORM} GOARCH=${ARCH} go build -o ${PLATFORM_OUTPUT_DIR}/${OUTPUT_NAME}

        # 创建压缩包
        echo "Creating ${PLATFORM} ${ARCH} release packages..."
        (cd "${PLATFORM_OUTPUT_DIR}" && tar -czvf "../${RELEASE_NAME}_${PLATFORM}_${ARCH}.tar.gz" ${OUTPUT_NAME})
        (cd "${PLATFORM_OUTPUT_DIR}" && zip -r "../${RELEASE_NAME}_${PLATFORM}_${ARCH}.zip" ${OUTPUT_NAME})
    done
done

# 清理临时文件
echo "Cleaning up..."
#rm -rf ${OUTPUT_DIR}

echo "Release process completed!"
