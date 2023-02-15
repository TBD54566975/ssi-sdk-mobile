#import "RNSsi.h"
#import <Ssi/Ssi.h>

@implementation RNSsi
RCT_EXPORT_MODULE()

RCT_REMAP_METHOD(generateDidKey,
                 ofType:(NSString*)key
                 withResolver:(RCTPromiseResolveBlock)resolve
                 withRejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error = [[NSError alloc] init];
    
    @try  {
        SsiDIDKeyWrapper *thing = SsiGenerateDIDKey(@"secp256k1", &error);
        resolve(thing.didKey);
    } @catch (NSException *exception) {
        //        reject(exception);
        //        reject(@"error", @"error description", error);
        
    }
}

// Don't compile this code when we build for the old architecture.
#ifdef RCT_NEW_ARCH_ENABLED
- (std::shared_ptr<facebook::react::TurboModule>)getTurboModule:
    (const facebook::react::ObjCTurboModule::InitParams &)params
{
    return std::make_shared<facebook::react::NativeSsiSpecJSI>(params);
}
#endif

@end
