
#ifdef RCT_NEW_ARCH_ENABLED
#import "RNSsiSpec.h"

@interface RNSsi : NSObject <NativeSsiSpec>
#else
#import <React/RCTBridgeModule.h>

@interface RNSsi : NSObject <RCTBridgeModule>
#endif

@end
