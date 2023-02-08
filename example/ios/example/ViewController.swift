//
//  ViewController.swift
//  example
//
//  Created by Tim Shamilov on 2/7/23.
//

import UIKit
import Identity

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view.
        var error: NSError? = NSError()
        
//        let supportedKeyTypes = IdentityGetSupportedKeyTypes();
        let did = IdentityGenerateDIDKey("RSA", &error);
        
        print(did?.didKey);
        if let unwrapped = did?.privateJSONWebKey {
            do {
                let json = try JSONSerialization.jsonObject(with: unwrapped);
                print(json)
            } catch {
                print("Error while parsing JSON: \(error)")
            }
        }
    }


}

