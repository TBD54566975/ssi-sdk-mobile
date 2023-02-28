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

// wip
RCT_REMAP_METHOD(createVerifiableCredential,
                 fromJSON:(NSDictionary*)json
                 withResolver:(RCTPromiseResolveBlock)resolve
                 withRejecter:(RCTPromiseRejectBlock)reject)
{
    SsiVerifiableCredentialMobile *verifiableCredential = [[SsiVerifiableCredentialMobile alloc] init];

    verifiableCredential.context = [json objectForKey:@"context"];
    verifiableCredential.id_ = [json objectForKey:@"id_"];
    verifiableCredential.type = [json objectForKey:@"type"];
    verifiableCredential.issuer = [json objectForKey:@"issuer"];
    verifiableCredential.issuanceDate = [json objectForKey:@"issuanceDate"];
    verifiableCredential.expirationDate = [json objectForKey:@"expirationDate"];
    verifiableCredential.credentialStatus = [json objectForKey:@"credentialStatus"];
    verifiableCredential.credentialSubject = [json objectForKey:@"credentialSubject"];
    
    // missing?
    // verifiableCredential.credentialSchema = [json objectForKey:@"credentialSchema"];
    // verifiableCredential.refreshService = [json objectForKey:@"refreshService"];
    
    verifiableCredential.termsOfUse = [json objectForKey:@"termsOfUse"];
    verifiableCredential.evidence = [json objectForKey:@"evidence"];
    verifiableCredential.proof = [json objectForKey:@"evidence"];
    
    // broken, not exported due to bug
    // verifiableCredential.toGoRepresentation
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
