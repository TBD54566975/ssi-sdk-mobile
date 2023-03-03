#import "RNSsi.h"
#import <Ssi/Ssi.h>

@implementation RNSsi
RCT_EXPORT_MODULE()

RCT_REMAP_METHOD(generateDidKey,
                 ofType:(NSString*)keyType
                 withResolver:(RCTPromiseResolveBlock)resolve
                 withRejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error = [[NSError alloc] init];
    
    @try  {
        SsiDIDKeyWrapper *thing = SsiGenerateDIDKey(keyType, &error);
        resolve(thing.didKey);
    } @catch (NSException *exception) {
        reject(@"Something weng wrong", @"wrong", error);
    }
}

RCT_REMAP_METHOD(expandDidKey,
                 forKey:(NSString*)didKey
                 resolver:(RCTPromiseResolveBlock)resolve
                 rejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error;
    NSData *bytes = SsiExpandDIDKey(didKey, &error);
    NSMutableDictionary *json = [NSJSONSerialization JSONObjectWithData:bytes options:NSJSONReadingMutableContainers error:&error];

    if (error == nil) {
        resolve(json);
    } else {
        reject(@"Something went wrong", @"wrong", error);
    }
}

RCT_REMAP_METHOD(createDidKey,
                 ofType:(NSString*)keyType
                 withPublicKey:(NSData*)publicKey
                 withResolver:(RCTPromiseResolveBlock)resolve
                 withRejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error = [[NSError alloc] init];
    
    @try  {
        NSString *thing = SsiCreateDIDKey(keyType, publicKey, &error);
    } @catch (NSException *exception) {
        reject(@"Something weng wrong", @"wrong", error);
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
