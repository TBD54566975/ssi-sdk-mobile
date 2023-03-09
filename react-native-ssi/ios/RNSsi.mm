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

RCT_REMAP_METHOD(signVerifiableCredentialJWT,
                 keyId:(NSString *)keyId
                 privateJwk:(NSDictionary *)privateJwk
                 verifiableCredential:(NSDictionary *)vc
                 resolver:(RCTPromiseResolveBlock)resolve
                 rejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error;
    NSData *privateJwkBytes = [NSJSONSerialization dataWithJSONObject:privateJwk options:NSJSONWritingPrettyPrinted error:&error];
    if (error != nil) {
        return reject(@"RNSsi", @"error serializing privateJwk bytes", error);
    }

    NSData *vcBytes = [NSJSONSerialization dataWithJSONObject:vc options:NSJSONWritingPrettyPrinted error:&error];
    if (error != nil) {
        return reject(@"RNSsi", @"error serializing vc bytes", error);
    }

    NSString *result = SsiSignVerifiableCredentialJWT(keyId, privateJwkBytes, vcBytes, &error);
    if (error == nil) {
        resolve(result);
    } else {
        reject(@"RNSsi", @"error signing vc", error);
    }
}

RCT_REMAP_METHOD(verifyVerifiableCredentialJWT,
                 keyId:(NSString *)keyId
                 publicJwk:(NSDictionary *)publicJwk
                 jwt:(NSString *)jwt
                 resolver:(RCTPromiseResolveBlock)resolve
                 rejecter:(RCTPromiseRejectBlock)reject)
{
    NSError *error;
    NSData *publicJwkBytes = [NSJSONSerialization dataWithJSONObject:publicJwk options:NSJSONWritingPrettyPrinted error:&error];
    if (error != nil) {
        return reject(@"RNssi", @"error serializing publicJwk bytes", error);
    }

    NSData *vcBytes = SsiVerifyVerifiableCredentialJWT(keyId, publicJwkBytes, jwt, &error);
    if (error != nil) {
        return reject(@"RNSsi", @"error verifying vc", error);
    }

    NSMutableDictionary *json = [NSJSONSerialization JSONObjectWithData:vcBytes options:NSJSONReadingMutableContainers error:&error];
    if (error == nil) {
        resolve(json);
    } else {
        reject(@"RNSsi", @"error serializing verified vc json", error);
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
