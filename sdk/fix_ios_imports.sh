#!/bin/bash

find Ssi.xcframework -name '*.h' | xargs sed -i '' -e 's~@import Foundation;~#import <Foundation/Foundation.h>~g';