#import "RNSsi.h"
#import <Ssi/Ssi.h>

@implementation RNSsi
RCT_EXPORT_MODULE()

RCT_REMAP_METHOD(generateDidKey,
                 ofType:(NSString*)keyType
                 withResolver:(RCTPromiseResolveBlock)resolve
                 withRejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error;
    NSData *bytes = SsiGenerateDIDKey(keyType, &error);
    if (error != nil) {
        return reject(@"RNSsi", @"error generating DID key", error);
    }

    NSMutableDictionary *json = [NSJSONSerialization JSONObjectWithData:bytes options:NSJSONReadingMutableContainers error:&error];
    if (error != nil) {
        reject(@"RNSsi", @"error serializing json", error);
    } else {
        resolve(json);
    }
}

RCT_REMAP_METHOD(expandDidKey,
                 forKey:(NSString*)didKey
                 resolver:(RCTPromiseResolveBlock)resolve
                 rejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error;
    NSData *bytes = SsiExpandDIDKey(didKey, &error);
    if (error != nil) {
        return reject(@"RNSsi", @"error expanding DID key", error);
    }

    NSMutableDictionary *json = [NSJSONSerialization JSONObjectWithData:bytes options:NSJSONReadingMutableContainers error:&error];
    if (error != nil) {
        reject(@"RNSsi", @"error serializing json", error);
    } else {
        resolve(json);
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
