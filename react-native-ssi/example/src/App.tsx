import React from 'react';

import {
  StyleSheet,
  View,
  Text,
  ScrollView,
  SafeAreaView,
  TouchableOpacity,
} from 'react-native';
import { generateDidKey, expandDidKey } from 'react-native-ssi';

export function App() {
  const [logDisplay, setLogDisplay] = React.useState('App Initialized. \n\n');
  const [didKey, setDidKey] = React.useState<string>('');
  // const [privateJwk, setPrivateJwk] = React.useState<Record<string, unknown>>();

  const addLogLine = (text: unknown) => {
    setLogDisplay(
      (previousLogs) => previousLogs + JSON.stringify(text) + '\n\n'
    );
  };

  return (
    <SafeAreaView style={styles.container}>
      <ScrollView showsVerticalScrollIndicator={false}>
        <View style={styles.container}>
          <>
            <TouchableOpacity
              style={styles.button}
              onPress={() => {
                generateDidKey('RSA').then((result) => {
                  setDidKey(result.didKey);
                  addLogLine(didKey);

                  // currently unused
                  // setPrivateJwk(result.privateJwk);
                  // addLogLine(privateJwk);
                });
              }}
            >
              <Text style={styles.buttonText}>Generate DID</Text>
            </TouchableOpacity>
            <TouchableOpacity
              disabled={didKey === ''}
              style={styles.button}
              onPress={() => {
                expandDidKey(didKey).then((result) => {
                  addLogLine(result);
                });
              }}
            >
              <Text style={styles.buttonText}>Expand DIDDoc</Text>
            </TouchableOpacity>
            <Text style={styles.logDisplay}>{logDisplay}</Text>
          </>
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    marginHorizontal: 16,
  },
  button: {
    backgroundColor: '#2196F3',
    borderRadius: 20,
    elevation: 2,
    padding: 8,
    marginVertical: 4,
    marginRight: 6,
  },
  buttonText: {
    color: 'white',
    fontWeight: 'bold',
    textAlign: 'center',
  },
  logDisplay: {
    marginTop: 8,
    marginBottom: 8,
    paddingHorizontal: 16,
  },
});
