#ifdef __cplusplus
#import "react-native-ssi.h"
#endif

#ifdef RCT_NEW_ARCH_ENABLED
#import "RNSsiSpec.h"

@interface Ssi : NSObject <NativeSsiSpec>
#else
#import <React/RCTBridgeModule.h>

@interface Ssi : NSObject <RCTBridgeModule>
#endif

@end
