import React, { Component } from 'react';
import {StyleSheet, View, FlatList, Text} from 'react-native';

export default class UserList extends Component {
    render() {
        const players = this.props.players;
        return (
            <View style={styles.container}>
                { players && players.length > 0 ?
                    (<FlatList
                        data={players}
                        keyExtractor={(item, index) => item.ID.toString()}
                        renderItem={({item}) => <Text style={styles.item}>{item.name}</Text>}
                    />) : (<Text>Ei pelaajia</Text>)

                }
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        paddingTop: 22
    },
    item: {
        padding: 6,
        fontSize: 16,
        height: 30,
    },
});
