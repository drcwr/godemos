cmake_minimum_required(VERSION 3.6)
project(green_c___sdk)

set(CMAKE_CXX_STANDARD 11)

set(SOURCE_FILES
        ApiBase.h
        Request.h
        DefaultProfile.h
        GreenService.h
        ImageSyncScanSample.cc
        utils.h
        Biz.h)

link_libraries("${PROJECT_SOURCE_DIR}/lib/libssl.a")
link_libraries("${PROJECT_SOURCE_DIR}/lib/libcrypto.a")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -lcurl")
add_executable(green_c___sdk ${SOURCE_FILES})
target_link_libraries(green_c___isdk curl)