export default class UserResponsesList extends Component {
    render() {
        const players = this.props.players;
        return (
            <View style={styles.container}>
                <FlatList
                    data={players}
                    renderItem={({item}) => <View key={item.player_name}><Text style={styles.item}>{item.player_name}</Text><Text>{item.score}</Text></View>}
                />
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
