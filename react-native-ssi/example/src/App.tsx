import * as React from 'react';

import {
  StyleSheet,
  View,
  Text,
  ScrollView,
  SafeAreaView,
  TouchableOpacity,
} from 'react-native';
import { generateDidKey } from 'react-native-ssi';

export default function App() {
  const [logDisplay, setLogDisplay] = React.useState('App Initialized. \n\n');

  const addLogLine = (text: string) => {
    setLogDisplay((previousLogs) => previousLogs + text + '\n\n');
  };

  return (
    <SafeAreaView style={styles.container}>
      <ScrollView showsVerticalScrollIndicator={false}>
        <View style={styles.container}>
          <>
            <TouchableOpacity
              style={styles.button}
              onPress={() => generateDidKey('secp256k1').then(addLogLine)}
            >
              <Text style={styles.buttonText}>Generate DID</Text>
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
  box: {
    width: 60,
    height: 60,
    marginVertical: 20,
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
