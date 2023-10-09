#import "CryptoExportPlugin.h"
#if __has_include(<crypto_export/crypto_export-Swift.h>)
#import <crypto_export/crypto_export-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "crypto_export-Swift.h"
#endif

@implementation CryptoExportPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftCryptoExportPlugin registerWithRegistrar:registrar];
}
@end
