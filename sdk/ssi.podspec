Pod::Spec.new do |s|
  s.name         = "ssi"
  s.version      = "0.0.1"
  s.summary      = "SDK for Self Sovereign Identity and Verifiable Credentials"
  s.homepage     = "https://github.com/TBD54566975/ssi-sdk-mobile.git"
  s.license      = "Apache"
  s.authors      = "Tim Shamilov", "Adam Mika", "Gabe Cohen", "Andres Uribe"
  s.source       = { :git => "https://github.com/TBD54566975/ssi-sdk-mobile.git", :tag => "#{s.version}" }
  s.vendored_frameworks = "Ssi.xcframework"
end
