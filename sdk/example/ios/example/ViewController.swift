//
//  ViewController.swift
//  example
//
//  Created by Tim Shamilov on 2/7/23.
//

import UIKit
import Ssi

class ViewController: UIViewController {
    
    @IBOutlet weak var LittleConsole: UILabel!
    @IBAction func CreateDidTapped(_ sender: Any) {
        createDID();
    }
    
    override func viewDidLoad() {
        super.viewDidLoad()
        addLog(text: "App initialized.");
    }
    
    func createDID() {
        // dont think we need this
        // let supportedKeyTypes = IdentityGetSupportedKeyTypes();
        
        var error: NSError? = NSError()
        let did = SsiGenerateDIDKey("secp256k1", &error);
        
        if let unwrapped = did?.privateJSONWebKey {
            do {
                let json = try JSONSerialization.jsonObject(with: unwrapped);
                addLog(text: "Did creation successful. Did is: \(did!.didKey)");
                addLog(text: "Did JSON is: \(json)");
            } catch {
                addLog(text: "Error while parsing JSON: \(error)")
            }
        }
    }
    
    
    // put private key in secure storage
    func storePrivateKey(key: String) throws {
        let tag = "com.tbd.example".data(using: .utf8)!
        let query: [String: Any] = [kSecClass as String: kSecClassKey,
                                       kSecAttrApplicationTag as String: tag,
                     kSecValueRef as String: key]
        
        let status = SecItemAdd(query as CFDictionary, nil)
        guard status == errSecSuccess else { throw KeychainError.failure }
    }
    
    // utils
    func addLog(text: String) {
        let currentLogs = self.LittleConsole.text ?? "";
        self.LittleConsole.text = currentLogs + text + "\n\n";
        print(text);
    }
    
    enum KeychainError: Error {
          case failure
      }
}

