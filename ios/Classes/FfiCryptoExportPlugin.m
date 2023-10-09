#import "FfiCryptoExportPlugin.h"
#if __has_include(<ffi_crypto_export/ffi_crypto_export-Swift.h>)
#import <ffi_crypto_export/ffi_crypto_export-Swift.h>
#else
// Support project import fallback if the generated compatibility header
// is not copied when this plugin is created as a library.
// https://forums.swift.org/t/swift-static-libraries-dont-copy-generated-objective-c-header/19816
#import "ffi_crypto_export-Swift.h"
#endif

@implementation FfiCryptoExportPlugin
+ (void)registerWithRegistrar:(NSObject<FlutterPluginRegistrar>*)registrar {
  [SwiftFfiCryptoExportPlugin registerWithRegistrar:registrar];
}
@end
