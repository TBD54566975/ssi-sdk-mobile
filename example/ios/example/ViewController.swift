//
//  ViewController.swift
//  example
//
//  Created by Tim Shamilov on 2/7/23.
//

import UIKit
import Identity

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
        let did = IdentityGenerateDIDKey("secp256k1", &error);
        
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
    
    // utils
    func addLog(text: String) {
        let currentLogs = self.LittleConsole.text ?? "";
        self.LittleConsole.text = currentLogs + text + "\n\n";
        print(text);
    }
}

